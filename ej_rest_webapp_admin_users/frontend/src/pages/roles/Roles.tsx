/* eslint-disable jsx-a11y/anchor-is-valid */
import axios from 'axios';
import {useEffect, useState} from 'react';
import Wrapper from "../../components/Wrapper";
import {Role} from "../../models/role";
import {Link} from "react-router-dom";

const Roles = () => {
    const [roles, setRoles] = useState([]);

    useEffect(() => {
        (
            async () => {
                const {data} = await axios.get('roles');
                setRoles(data);
            }
        )();
    }, []);

    const del = async (id: number) => {
        if (window.confirm('You want to delete this record?')) {
            await axios.delete(`delete-role/${id}`);

            setRoles(roles.filter((r: Role) => r.role_id !== id));
        }
    }

    return (
        <Wrapper>
            <div className="pt-3 pb-2 mb-3 border-bottom">
                <Link to="/roles/create" className="btn btn-sm btn-outline-secondary">Create</Link>
            </div>
            <div className="table-responsive">
                <table className="table table-striped table-sm">
                    <thead>
                    <tr>
                        <th>#</th>
                        <th>Name</th>
                        <th>Action</th>
                    </tr>
                    </thead>
                    <tbody>
                    {roles.map((role: Role) => {
                        return (
                            <tr key={role.role_id}>
                                <td>{role.role_id}</td>
                                <td>{role.name}</td>
                                <td>
                                    <div className="btn-group mr-2">
                                        <Link 
                                        to={`/roles/${role.role_id}/edit`}
                                        className="btn btn-sm btn-outline-secondary"
                                        >Edit</Link>
                                    </div>
                                    &nbsp;
                                    <div className="btn-group mr-2"> 
                                        <a href="#"
                                        className="btn btn-sm btn-outline-secondary"
                                        onClick={() => del(role.role_id)}
                                        >Delete</a>
                                    </div>
                                </td>
                            </tr>
                        )
                    })}
                    </tbody>
                </table>
            </div>
        </Wrapper>
    );
};

export default Roles;