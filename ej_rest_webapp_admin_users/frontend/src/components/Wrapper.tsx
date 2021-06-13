import axios from "axios"
import {useEffect, useState} from "react"
import {Redirect} from "react-router-dom"
import Menu from "./Menu"
import Nav from "./Nav"

const Wrapper = (props: any) => {
    const [redirect, setRedirect] = useState(false);
    useEffect(() => {
        (async () => {
            try {
                await axios.get('user'); 
            } catch (error) {
                setRedirect(true);
            }
            
        })()
    }, [])
    
    if (redirect) {
        return <Redirect to={'/login'}/>
    }    

    return (
    <>
    <Nav />
    <div className="container-fluid">
        <div className="row">
            <Menu />
            <main className="col-md-9 ms-sm-auto col-lg-10 px-md-4">
                {props.children}
            </main>
        </div>
    </div>
    </>
    )
}

export default Wrapper;