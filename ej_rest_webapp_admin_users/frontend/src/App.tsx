import Menu from "./components/Menu";
import Nav from "./components/Nav";
import './css/App.css';

function App() {
  return (
  <div className="App">
    <Nav />
    <div className="container-fluid">
      <div className="row">
        <Menu />
        <main className="col-md-9 ms-sm-auto col-lg-10 px-md-4">
          <div className="table-responsive">
            <table className="table table-striped table-sm">
              <thead>
                <tr>
                  <th>#</th>
                  <th>Header</th>
                  <th>Header</th>
                  <th>Header</th>
                  <th>Header</th>
                </tr>
              </thead>
              <tbody>
                <tr>
                  <td>1,001</td>
                  <td>random</td>
                  <td>data</td>
                  <td>placeholder</td>
                  <td>text</td>
                </tr>
              </tbody>
            </table>
          </div>
        </main>
      </div>
    </div>
  </div>
  );
}

export default App;