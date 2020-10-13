import React from 'react';
import { API_HOST } from "../../../utils/config";
import AvatarNotFound from "../../../assets/png/avatar-no-found.png";
import "./Media.scss";

export default function Media(props) {
    const {user}=props;
    const bannerUrl=user?.banner? `${API_HOST}/obtenerbanner?id=${user.id}`:null;
    const avatarUrl=user?.avatar? `${API_HOST}/obteneravatar?id=${user.id}`:AvatarNotFound;
    return (
        <div className="banner-avatar" style={{ backgroundImage: `url('${bannerUrl}')` }}>
            <div className="avatar" style={{ backgroundImage: `url('${avatarUrl}')` }}></div>
        </div>
    )
}
