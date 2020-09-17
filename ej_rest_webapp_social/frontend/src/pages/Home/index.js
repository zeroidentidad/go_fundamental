import React from 'react'
import "./Home.scss";
import {BasicLayout} from "../../layout"

export default function Home(props) {
    const { setRefreshCheckLogin }=props;
    return (
    <BasicLayout className="home" setRefreshCheckLogin={setRefreshCheckLogin}>
        <h1>En home</h1>
    </BasicLayout>
    )
}
