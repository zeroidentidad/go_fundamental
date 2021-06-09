import {BrowserRouter, Route} from "react-router-dom";
import './css/App.css';
import Dashboard from "./pages/Dashboard";
import Register from "./pages/Register";
import Users from "./pages/Users";

function App() {
  return (
  <div className="App">
    <BrowserRouter>
      <Route path={'/'} exact component={Dashboard} />
      <Route path={'/users'} component={Users} />
      <Route path={'/register'} component={Register} />    
    </BrowserRouter>
  </div>
  );
}

export default App;