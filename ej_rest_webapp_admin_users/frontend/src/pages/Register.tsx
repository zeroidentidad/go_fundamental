import { Component, SyntheticEvent } from 'react';
import '../css/Register.css';

export default class Register extends Component {
    first_name = '';
    last_name = '';
    email = '';
    password = '';
    password_confirm = '';

    submit = async (e: SyntheticEvent) => {
        e.preventDefault();

        console.log([
            this.first_name,
            this.last_name,
            this.email,
            this.password,
            this.password_confirm
        ])
    }    

    render() {
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
