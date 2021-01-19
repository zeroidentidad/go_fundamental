import React from 'react';
import {BasicLayout} from "../../layout";

import "./Users.scss";

export default function Users(props) {
    const {setRefreshCheckLogin} = props;
    return (
        <BasicLayout className="users" setRefreshCheckLogin={setRefreshCheckLogin}>
            <div>
                users
            </div>
        </BasicLayout>
    )
}
