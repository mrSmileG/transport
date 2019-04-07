import React, { Component } from 'react';
import { Table } from 'semantic-ui-react';

class SolutionTable extends Component {
    static async getInitialProps(props) {
        const {
            data
        } = props.query;

        return { data };
    }

    getTotalPrice = () => {
        var sum = 0;
        this.props.data.map(item => item.map(i => (
                sum+=i.count*i.price
        )))

        return sum;
    }

    render() {
        const {
            Cell,
            Row,
            HeaderCell,
            Body,
            Header
        } = Table;

        const {
            data
        } = this.props;

        return data.length === 0 ? "" : (
            <Table>
                <Header>
                    <Row>
                        <HeaderCell>C(n)\P(n)</HeaderCell>
                        {data.slice(0, 1).map(item =>                         
                            (item.map((i, index) => (
                                <HeaderCell id={i}>{"P" + (index+1)}</HeaderCell>                                
                            )))
                        )}
                    </Row>
                </Header>
                <Body>
                {data.map((item, index) => 
                    <Row>
                        <Cell>{"C" + (index+1)}</Cell>
                        {(item.map(i => (
                            <Cell>{i.count}</Cell>
                        )))}
                    </Row>

                )}
                    <Row positive>
                        <Cell>Total Price:</Cell>
                        <Cell>{this.getTotalPrice() + "$"}</Cell>
                    </Row>
                </Body>
            </Table>
        );
    }
}

export default SolutionTable;