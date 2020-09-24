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
