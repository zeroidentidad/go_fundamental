import React, {useState, useEffect} from 'react';
import {withRouter} from "react-router-dom";
import { toast } from "react-toastify";
import {BasicLayout} from "../../layout";
import Media from "../../components/User/Media";
import InfoUser from "../../components/User/InfoUser";
import ListTweets from "../../components/ListTweets";
import {getUserApi} from "../../api/user";
import { getUserTweetsApi } from "../../api/tweet";
import useAuth from "../../hooks/useAuth";

import "./User.scss";

function User(props) {
    const {match, setRefreshCheckLogin}=props;
    const {params}=match;
    const [user, setUser] = useState(null);
    const [tweets, setTweets] = useState(null);
    const loggedUser=useAuth();

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

    useEffect(() => {
        getUserTweetsApi(params.id, 1)
        .then((response) => {
            setTweets(response);
        })
        .catch(() => {
            setTweets([]);
        });
    }, [params]);    

    return (
        <BasicLayout className="user" setRefreshCheckLogin={setRefreshCheckLogin}>
            <div className="user__title">
                <h2>{user?`${user.nombre} ${user.apellidos}`:'El usuario no existe'}</h2>
            </div>
            <Media user={user} loggedUser={loggedUser}/>
            <InfoUser user={user}/>
            <div className="user__tweets">
                <h3>Tweets</h3>
                {tweets && <ListTweets tweets={tweets} />}
            </div>
        </BasicLayout>
    )
}

export default withRouter(User)