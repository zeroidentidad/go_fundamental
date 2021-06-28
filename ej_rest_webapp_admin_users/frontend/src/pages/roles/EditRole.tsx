import axios from "axios";
import {SyntheticEvent, useEffect, useState} from "react";
import {Redirect} from "react-router-dom";
import Wrapper from "../../components/Wrapper";
import {Permission} from "../../models/permission";

const EditRole = (props: any) => {
    const [permissions, setPermissions] = useState([]);
    const [selected, setSelected] = useState([] as string[]);
    const [name, setName] = useState('');
    const [redirect, setRedirect] = useState(false);    

    useEffect(() => {
        (
            async () => {
                const {data} = await axios.get('permissions');
                setPermissions(data);
                
                const response = await axios.get(`get-role/${props.match.params.id}`);
                setName(response.data.name);
                setSelected(response.data.permissions.map((p: Permission) => String(p.permission_id)));
            }
        )();
    }, [props.match.params.id]);

    const check = (id: string) => {
        if (selected.some(s => s === id)) {
            setSelected(selected.filter(s => s !== id));
            return;
        }
        setSelected([...selected, id]);
    }    

    const submit = async (e: SyntheticEvent) => {
        e.preventDefault();

        let role_id = Number(props.match.params.id)
        const response = await axios.put('update-role', {
            role_id,
            name,
            permissions: selected
        });

        if(response.status===201){
            setRedirect(true)
        }
    }

    if (redirect) {
        return <Redirect to="/roles"/>;
    }    
    
    return (
        <Wrapper>
            <form onSubmit={submit}>
                <div className="mb-3 mt-3 row">
                    <label className="col-sm-2 col-form-label">Name</label>
                    <div className="col-sm-10">
                        <input className="form-control" defaultValue={name} onChange={e => setName(e.target.value)}/>
                    </div>
                </div>
                <div className="mb-3 row">
                    <label className="col-sm-2 col-form-label">Permissions</label>
                    <div className="col-sm-10">
                        {permissions.map((p: Permission) => {
                            return (
                                <div className="form-check form-check-inline col-3" key={p.permission_id}>
                                    <input
                                    className="form-check-input" type="checkbox"
                                    value={p.permission_id}
                                    checked={selected.some(s => String(s) === String(p.permission_id))}
                                    onChange={() => check(String(p.permission_id))}
                                    />
                                    <label className="form-check-label">{p.name}</label>
                                </div>
                            )
                        })}
                    </div>
                </div>

                <button className="btn btn-outline-secondary">Save</button>
            </form>
        </Wrapper>
    )
}

export default EditRole;
