import { Component } from 'react'
import Wrapper from "../components/Wrapper"

export default class Users extends Component {
    render() {
        return (
            <Wrapper>
            <div className="table-responsive">
                <table className="table table-striped table-sm">
                <thead>
                    <tr>
                    <th>#</th>
                    <th>Header</th>
                    <th>Header</th>
                    <th>Header</th>
                    <th>Header</th>
                    </tr>
                </thead>
                <tbody>
                    <tr>
                    <td>1,001</td>
                    <td>random</td>
                    <td>data</td>
                    <td>placeholder</td>
                    <td>text</td>
                    </tr>
                </tbody>
                </table>
            </div>
            </Wrapper>
        )
    }
}
