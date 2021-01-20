import React, { useState, useEffect } from 'react';
import { withRouter } from "react-router-dom";
import { Spinner, ButtonGroup, Button } from "react-bootstrap";
import queryString from "query-string";
import {BasicLayout} from "../../layout";
import { getFollowsApi } from "../../api/follow";

import "./Users.scss";

function Users(props) {
    const {setRefreshCheckLogin, location, history} = props;
    const [users, setUsers] = useState(null);
    const params = useUsersQuery(location);
    const [btnLoading, setBtnLoading] = useState(false);

  useEffect(() => {
    getFollowsApi(queryString.stringify(params))
      .then((response) => {
        // eslint-disable-next-line eqeqeq
        if (params.page == 1) {
          if (isEmpty(response)) {
            setUsers([]);
          } else {
            setUsers(response);
          }
        } else {
          if (!response) {
            setBtnLoading(0);
          } else {
            setUsers([...users, ...response]);
            setBtnLoading(false);
          }
        }
      })
      .catch(() => {
        setUsers([]);
      });
    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [location]);


    return (
        <BasicLayout className="users" setRefreshCheckLogin={setRefreshCheckLogin}>
            <div className="users__title">
                <h2>Usuarios</h2>
                <input
                type="text"
                placeholder="Buscar usuario..."
                />
            </div>
            <ButtonGroup className="users__options">
                <Button
                >
                Siguiendo
                </Button>
                <Button
                >
                Nuevos
                </Button>
            </ButtonGroup>            
        </BasicLayout>
    )
}

function useUsersQuery(location) {
  const { page = 1, type = "follow", search="" } = queryString.parse(
    location.search
  );

  return { page, type, search };
}

export default withRouter(Users);