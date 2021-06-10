import axios from "axios";
import {SyntheticEvent, useState} from "react";
import {Redirect} from "react-router";

const Login = () => {
    const [email, setEmail] = useState("");
    const [password, setPassword] = useState("");
    const [redirect, setRedirect] = useState(false);

    const submit = async (e: SyntheticEvent) => {
        e.preventDefault();

         const response = await axios.post('login', {
            email,
            password
        });

        if(response.status===201){
            setRedirect(true)
        }
    }

    if (redirect) {
        return <Redirect to={'/'}/>   
    }    

    return (
            <main className="form-signin">
            <form onSubmit={submit}>
                <h1 className="h3 mb-3 fw-normal">Please sign in</h1>
                <input type="email" className="form-control" placeholder="name@example.com"
                onChange={e=>setEmail(e.target.value)}
                />
                <br/>
                <input type="password" className="form-control" placeholder="Password" 
                onChange={e=>setPassword(e.target.value)}
                />
                <br/>
                <button className="w-100 btn btn-lg btn-primary" type="submit">Submit</button>
            </form>
            </main>
    )
}

export default Login;