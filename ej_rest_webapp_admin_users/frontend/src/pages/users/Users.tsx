import axios from "axios";
import {useEffect, useState} from "react";
import Paginator from "../../components/Paginator";
import Wrapper from "../../components/Wrapper";
import {User} from "../../models/user";

const Users = () => {
    const [users, setUsers] = useState([])
    const [page, setPage] = useState(1);
    const [lastPage, setLastPage] = useState(0);

    useEffect(() => {
        (
            async () => {
                const {data} = await axios.get(`users?page=${page}`);
                setUsers(data.data)
                setLastPage(data.meta.last_page);
            }
        )()
    }, [page])

    return (
            <Wrapper>
            <div className="table-responsive">
                <table className="table table-striped table-sm">
                <thead>
                    <tr>
                    <th>#</th>
                    <th>Name</th>
                    <th>Email</th>
                    <th>Role</th>
                    <th>Action</th>
                    </tr>
                </thead>
                <tbody>
                    {users.map((user: User)=>{
                    return(
                    <tr key={user.user_id}>
                    <td>{user.user_id}</td>
                    <td>{user.first_name} {user.last_name}</td>
                    <td>{user.email}</td>
                    <td>{user.role.name}</td>
                    <td>action</td>
                    </tr>                         
                    )                       
                    })}
                </tbody>
                </table>
            </div>
            <Paginator page={page} lastPage={lastPage} pageChanged={setPage}/>           
            </Wrapper>
    );
}

export default Users;
