import React, {useState, useEffect} from 'react';
import {withRouter} from "react-router-dom";
import { toast } from "react-toastify";
import { Button, Spinner } from "react-bootstrap";
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
    const [page, setPage] = useState(1);
    const [loadingTweets, setLoadingTweets] = useState(false);
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
    
    const moreData = () => {
        const pageTemp = page + 1;
        setLoadingTweets(true);

        getUserTweetsApi(params.id, pageTemp).then((response) => {
        if (!response) {
            setLoadingTweets(0);
        } else {
            setTweets([...tweets, ...response]);
            setPage(pageTemp);
            setLoadingTweets(false);
        }
        });
    };    

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
                <Button onClick={moreData}>
                {!loadingTweets ? (
                    loadingTweets !== 0 && "Obtener más Tuits"
                ) : (
                    <Spinner
                    as="span"
                    animation="grow"
                    size="sm"
                    role="status"
                    arian-hidden="true"
                    />
                )}
                </Button>                
            </div>
        </BasicLayout>
    )
}

export default withRouter(User)