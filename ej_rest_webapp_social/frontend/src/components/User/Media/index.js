import React from 'react';
import { API_HOST } from "../../../utils/config";

import "./Media.scss";

export default function Media(props) {
    const {user}=props;
    const bannerUrl=user?.banner? `${API_HOST}/obtenerbanner?id=${user.id}`:null;
    return (
        <div className="banner-avatar" style={{ backgroundImage: `url('${bannerUrl}')` }}>
            <div className="avatar"></div>
        </div>
    )
}
