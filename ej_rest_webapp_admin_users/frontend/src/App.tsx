import {BrowserRouter, Route} from "react-router-dom";
import './css/App.css';
import Dashboard from "./pages/Dashboard";
import Login from "./pages/Login";
import Register from "./pages/Register";
import Users from "./pages/users/Users";

function App() {
  return (
  <div className="App">
    <BrowserRouter>
      <Route path={'/'} exact component={Dashboard} />
      <Route path={'/users'} component={Users} />
      <Route path={'/register'} component={Register} />    
      <Route path={'/login'} component={Login} />    
    </BrowserRouter>
  </div>
  );
}

export default App;