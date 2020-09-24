import React, {useState, useEffect} from 'react';
import {Button, Spinner} from "react-bootstrap";
import {withRouter} from "react-router-dom";
import {BasicLayout} from "../../layout";
import {getUserApi} from "../../api/user";
import { toast } from "react-toastify";

import "./User.scss";

function User(props) {
    const {match}=props;
    const {params}=match;
    const [user, setUser] = useState(null);

    useEffect(() => {

        getUserApi(params.id)
        .then(res => {
            if (!res) toast.error("El usuario no existe");
            setUser(res);
        })
        .catch(() => {
            toast.error("El usuario no existe");
        });

    }, [params])

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

export default withRouter(User)