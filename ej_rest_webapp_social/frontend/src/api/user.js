import { API_HOST } from "../utils/config";
import { getTokenApi } from "./auth";

export function getUserApi(id) {
    const url=`${API_HOST}/verperfil?id=${id}`;

    const params={
        headers: {
            "Content-Type": "application/json",
            Authorization: `Bearer ${getTokenApi()}`
        }
    };

    return fetch(url, params)
        .then(response => {
            // eslint-disable-next-line no-throw-literal
            if (response.status>=400) throw null;
            return response.json();
        })
        .then(result => {
            return result;
        })
        .catch(err => {
            return err;
        });
}

export function uploadBannerApi(file) {
    const url = `${API_HOST}/subirbanner`;
    const formData = new FormData();
    formData.append("banner", file);
    
    const params = {
        method: "POST",
        headers: {
            Authorization: `Bearer ${getTokenApi()}`
        },
        body: formData
    };
    
    return fetch(url, params)
    .then(response => {
        return response.json();
    })
    .then(result => {
        return result;
    })
    .catch(err => {
        return err;
    });
}

export function uploadAvatarApi(file) {
    const url = `${API_HOST}/subiravatar`;
    const formData = new FormData();
    formData.append("avatar", file);
    
    const params = {
        method: "POST",
        headers: {
            Authorization: `Bearer ${getTokenApi()}`
        },
        body: formData
    };
    
    return fetch(url, params)
    .then(response => {
        return response.json();
    })
    .then(result => {
        return result;
    })
    .catch(err => {
        return err;
    });
}
