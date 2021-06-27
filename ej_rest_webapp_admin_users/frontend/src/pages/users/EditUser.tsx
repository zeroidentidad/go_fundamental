import axios from 'axios';
import {SyntheticEvent, useEffect, useState} from 'react';
import {Redirect} from 'react-router-dom';
import Wrapper from "../../components/Wrapper";
import {Role} from "../../models/role";

const EditUser = (props: any) => {
    const [first_name, setFirstName] = useState('');
    const [last_name, setLastName] = useState('');
    const [email, setEmail] = useState('');
    const [password, setPassword] = useState('');
    const [role_id, setRoleId] = useState(0);
    const [roles, setRoles] = useState([]);
    const [redirect, setRedirect] = useState(false);

    useEffect(() => {
        (
            async () => {
                const {data} = await axios.get('roles');
                setRoles(data);

                const res = await axios.get(`get-user/${props.match.params.id}`);

                setFirstName(res.data.first_name);
                setLastName(res.data.last_name);
                setEmail(res.data.email);
                setRoleId(Number(res.data.role.role_id));                
            }
        )()
    }, [props.match.params.id]);

    const submit = async (e: SyntheticEvent) => {
        e.preventDefault();

        if (role_id===0) {
            alert("Select a role for user")
        } else {
            let user_id = Number(props.match.params.id)
            const response = await axios.put('update-user', {
                user_id,
                first_name,
                last_name,
                email,
                password,
                role_id
            });

            if(response.status===201){
                setRedirect(true)
            }
        }
    }

    if (redirect) {
        return <Redirect to="/users"/>
    }

    return (
        <Wrapper>
            <form onSubmit={submit}>
                <div className="mb-3">
                    <label>First Name</label>
                    <input className="form-control" defaultValue={first_name} onChange={e => setFirstName(e.target.value)}/>
                </div>
                <div className="mb-3">
                    <label>Last Name</label>
                    <input className="form-control" defaultValue={last_name} onChange={e => setLastName(e.target.value)}/>
                </div>
                <div className="mb-3">
                    <label>Email</label>
                    <input type="email" className="form-control" defaultValue={email} onChange={e => setEmail(e.target.value)}/>
                </div>
                <div className="mb-3">
                    <label>Password</label>
                    <input type="password" className="form-control" onChange={e => setPassword(e.target.value)}/>
                </div>
                <div className="mb-3">
                    <label>Role</label>
                    <select className="form-control" value={role_id} onChange={e => setRoleId(Number(e.target.value))}>
                        <option key={0} value={0}>Select role</option>
                        {roles.map((r: Role) => {
                            return (
                                <option key={r.role_id} value={r.role_id}>{r.name}</option>
                            )
                        })}
                    </select>
                </div>
                <button className="btn btn-outline-secondary">Save</button>
            </form>
        </Wrapper>
    )
}

export default EditUser;