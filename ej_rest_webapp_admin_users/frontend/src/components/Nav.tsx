import axios from "axios";
import {useEffect, useState} from "react";
import {NavLink} from "react-router-dom";
import {User} from "../models/user";

const Nav = () => {
    const [user, setUser] = useState(new User());
    useEffect(() => {
        (async () => {
            const {data} = await axios.get('user');
            setUser(new User(
                data.id,
                data.first_name,
                data.last_name,
                data.email,
                data.role
            ));
        })()
    }, [])

    const logout = async() => {
        await axios.post('logout', {});
    }

    return (
        <nav className="navbar navbar-dark sticky-top bg-dark flex-md-nowrap p-0 shadow">
            <a className="navbar-brand col-md-3 col-lg-2 mr-0 px-3" href="/">Company name</a>

            <ul className="my-2 my-md-0 mr-md-3">
                <NavLink to="/profile" className="p-2 text-white text-decoration-none">
                    {user.name}
                </NavLink>
                <NavLink to="/login" className="p-2 text-white text-decoration-none"
                onClick={logout}
                >Sign out
                </NavLink>
            </ul>
        </nav>
    )
}

export default Nav;