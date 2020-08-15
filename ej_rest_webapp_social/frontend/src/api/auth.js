import { API_HOST } from "../utils/config";

export function signUpApi(user) {
    const url=`${API_HOST}/registro`;
    console.log(user);
    console.log(url);
}