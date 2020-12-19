import React, {useState, useEffect} from 'react';
import {Button} from "react-bootstrap";
import { API_HOST } from "../../../utils/config";
import { checkFollowApi } from "../../../api/follow";
import ConfigModal from "../../Modales/ConfigModal";
import EditUserForm from "../EditUserForm";
import AvatarNotFound from "../../../assets/png/avatar-no-found.png";
import "./Media.scss";

export default function Media(props) {
    const {user, loggedUser}=props;
    const [showModal, setShowModal] = useState(false);
    const [following, setFollowing] = useState(null);
    const bannerUrl=user?.banner? `${API_HOST}/obtenerbanner?id=${user.id}`:null;
    const avatarUrl=user?.avatar? `${API_HOST}/obteneravatar?id=${user.id}`:AvatarNotFound;

    useEffect(() => {
        if (user) {
        checkFollowApi(user?.id).then(response => {
            if (response?.status) {
            setFollowing(true);
            } else {
            setFollowing(false);
            }
        });
        }

    }, [user]);    

    return (
        <div className="banner-avatar" style={{ backgroundImage: `url('${bannerUrl}')` }}>
            <div className="avatar" style={{ backgroundImage: `url('${avatarUrl}')` }}></div>
            {user&&(
                <div className="options">
                    {loggedUser._id===user.id&&(
                        <Button onClick={() => setShowModal(true)}>Editar perfil</Button>
                    )}

                    {loggedUser._id!==user.id && following !== null && (following ? (
                    <Button onClick={() => alert('siguiendo')}>
                        <span>Siguiendo</span>
                    </Button>
                    ) : (
                    <Button onClick={() => alert('seguir')}>
                        <span>Seguir</span>
                    </Button>
                    ))}
                </div>
            )}
            <ConfigModal show={showModal} setShow={setShowModal} title="Editar perfil">
                <EditUserForm user={user} setShowModal={setShowModal} />
            </ConfigModal>            
        </div>
    )
}
