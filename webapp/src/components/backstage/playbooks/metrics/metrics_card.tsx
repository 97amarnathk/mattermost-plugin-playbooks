// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

import React from 'react';
import styled from 'styled-components';
import {Duration} from 'luxon';
import {useIntl} from 'react-intl';

import {MetricType, PlaybookWithChecklist} from 'src/types/playbook';
import {HorizontalSpacer} from 'src/components/backstage/styles';
import {PlaybookStats} from 'src/types/stats';
import {formatDuration} from 'src/components/formatted_duration';
import BarGraph from 'src/components/backstage/playbooks/bar_graph';

interface Props {
    playbook: PlaybookWithChecklist;
    playbookStats: PlaybookStats;
    index: number;
}

const MetricsCard = ({playbook, playbookStats, index}: Props) => {
    const {formatMessage} = useIntl();
    const stats = makeCardStats(playbook, playbookStats, index);
    const transformFn = playbook.metrics[index].type === MetricType.Duration ?
        (val: number) => formatDuration(Duration.fromMillis(val)) : (val: number) => val;
    const valueTransformFn = playbook.metrics[index].type === MetricType.Duration ?
        (val: number) => formatDuration(Duration.fromMillis(val), 'narrow', 'truncate') : (val: number) => val;

    return (
        <Container>
            <Card>
                <SummaryCardInner>
                    <Cell>
                        <Title>{formatMessage({defaultMessage: 'Average value'})}</Title>
                        <Value>{transformFn(stats.average)}</Value>
                    </Cell>
                    <Cell>
                        <Title>{formatMessage({defaultMessage: '10-run average value'})}</Title>
                        <Row>
                            <Value>{transformFn(stats.rolling_average)}</Value>
                            {percentageChange(stats.rolling_average_change)}
                        </Row>
                    </Cell>
                    <Cell>
                        <Title>{formatMessage({defaultMessage: 'Value range'})}</Title>
                        <Value>
                            {valueTransformFn(stats.value_range[0])}
                            <ValueTo>{' ' + formatMessage({defaultMessage: 'to'}) + ' '}</ValueTo>
                            {valueTransformFn(stats.value_range[1])}
                        </Value>
                    </Cell>
                    <Cell>
                        {
                            stats.target &&
                            <>
                                <Title>{formatMessage({defaultMessage: 'Target value'})}</Title>
                                <Value>{transformFn(stats.target)}</Value>
                            </>
                        }
                    </Cell>
                </SummaryCardInner>
            </Card>
            <HorizontalSpacer size={16}/>
            <Card>
                <BarGraph
                    title={playbook.metrics[index].title + ' ' + formatMessage({defaultMessage: 'per run over the last 10 runs'})}
                    labels={['1', '2', '3', '4', '5', '6', '7', '8', '9', '10']}
                    data={stats.rolling_values}
                    color={'--center-channel-color-48'}
                    yAxesTicksCallback={(val, idx) => (idx % 2 === 0 ? '' : valueTransformFn(val).toString())}
                    tooltipTitleCallback={(label) => `Run #${label}`}
                    tooltipLabelCallback={(val) => transformFn(val).toString()}
                />
            </Card>
        </Container>
    );
};

interface MetricsCardStats {
    average: number;
    rolling_average: number;
    rolling_average_change: number;
    value_range: number[];
    rolling_values: number[];
    target?: number;
}

const makeCardStats = (playbook: PlaybookWithChecklist, stats: PlaybookStats, idx: number) => {
    return {
        average: stats.metric_overall_average[idx],
        rolling_average: stats.metric_rolling_average[idx],
        rolling_average_change: stats.metric_rolling_average_change[idx],
        value_range: stats.metric_value_range[idx],
        rolling_values: stats.metric_rolling_values[idx],
        target: playbook.metrics[idx].target,
    } as MetricsCardStats;
};

const Container = styled.div`
    display: flex;
`;

const Card = styled.div`
    border: 1px solid rgba(var(--center-channel-color-rgb), 0.04);
    box-shadow: 0 2px 3px rgba(var(--center-channel-color-rgb), 0.08);
    border-radius: 4px;
    background: var(--center-channel-bg);
    width: 533px;
`;

const SummaryCardInner = styled.div`
    display: grid;
    grid-template-columns: 1fr 1fr;
    grid-template-rows: 1fr 1fr;
    grid-gap: 24px;
    padding: 24px;
`;

const Cell = styled.div`
    display: flex;
    flex-direction: column;
`;

const Title = styled.div`
    margin-bottom: 4px;
    font-weight: 600;
    font-size: 14px;
    line-height: 20px;
    color: rgba(var(--center-channel-color-rgb), 0.72);
`;

const Value = styled.div`
    font-size: 20px;
    line-height: 24px;
    font-weight: 600;
`;

const ValueTo = styled.span`
    font-weight: 400;
`;

const Row = styled.div`
    display: flex;
    align-items: center;
`;

const percentageChange = (change: number) => {
    if (change >= 99999999 || change === 0) {
        return null;
    }
    const changeSymbol = (change > 0) ? 'icon-arrow-up' : 'icon-arrow-down';

    return (
        <PercentageChange>
            <i className={'icon ' + changeSymbol}/>
            {change + '%'}
        </PercentageChange>
    );
};

const PercentageChange = styled.div`
    margin-left: 12px;
    display: flex;
    flex-direction: row;
    border-radius: 10px;
    background-color: rgba(var(--online-indicator-rgb), 0.08);
    color: var(--online-indicator);
    font-size: 10px;
    line-height: 15px;

    > i {
        font-size: 12px;
    }
`;

export default MetricsCard;
