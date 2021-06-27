import {BrowserRouter, Route} from "react-router-dom";
import './css/App.css';
import Dashboard from "./pages/Dashboard";
import Login from "./pages/Login";
import Register from "./pages/Register";
import CreateUser from "./pages/users/CreateUser";
import Users from "./pages/users/Users";

function App() {
  return (
  <div className="App">
    <BrowserRouter>
      <Route path={'/'} exact component={Dashboard} />
      <Route path={'/register'} component={Register} />    
      <Route path={'/login'} component={Login} />
      <Route path={'/users'} exact component={Users} /> 
      <Route path={'/users/create'} component={CreateUser} /> 
    </BrowserRouter>
  </div>
  );
}

export default App;