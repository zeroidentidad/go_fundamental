import React from 'react';
import {Button, Spinner} from "react-bootstrap";
import {BasicLayout} from "../../layout";

import "./User.scss";

export default function User() {
    return (
        <BasicLayout className="user">
            <div className="user__title">
                <h2>UserName</h2>
            </div>
            <div>Banner</div>
            <div>User Info</div>
            <div className="user__tweets">List Tweets</div>
        </BasicLayout>
    )
}
