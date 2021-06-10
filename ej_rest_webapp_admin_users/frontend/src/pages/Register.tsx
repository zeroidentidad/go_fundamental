import { Component, SyntheticEvent } from 'react';
import axios from 'axios';
import '../css/Register.css';
import {Redirect} from "react-router";

export default class Register extends Component {
    first_name = '';
    last_name = '';
    email = '';
    password = '';
    password_confirm = '';
    state = {
        redirect: false
    }    

    submit = async (e: SyntheticEvent) => {
        e.preventDefault();

         const response = await axios.post('register', {
            first_name: this.first_name,
            last_name: this.last_name,
            email: this.email,
            password: this.password,
            password_confirm: this.password_confirm
        });

        if(response.status===201){
            this.setState({redirect: true})
        }
    }    

    render() {
        if (this.state.redirect) {
            return <Redirect to={'/login'}/>   
        }
        return (
            <main className="form-signin">
            <form onSubmit={this.submit}>
                <h1 className="h3 mb-3 fw-normal">Please register</h1>
                <input type="text" className="form-control" placeholder="First name"
                onChange={e => this.first_name = e.target.value} 
                />
                <br/>
                <input type="text" className="form-control" placeholder="Last name" 
                onChange={e => this.last_name = e.target.value}
                />
                <br/>
                <input type="email" className="form-control" placeholder="name@example.com" 
                onChange={e => this.email = e.target.value}
                />
                <br/>
                <input type="password" className="form-control" placeholder="Password"
                onChange={e => this.password = e.target.value}
                />
                <br/>
                <input type="password" className="form-control" placeholder="Password confirm" 
                onChange={e => this.password_confirm = e.target.value}
                />
                <br/>
                <button className="w-100 btn btn-lg btn-primary" type="submit">Submit</button>
            </form>
            </main>
        )
    }
}
