import './App.css';
import {useEffect} from "react";
import {Outlet, useNavigate} from "react-router-dom";
import CustomFooter from "./common/components/CustomFooter";
import Spinner from "./common/components/Spinner";
import {useAuth} from "./modules/auth/hooks/useAuth";
import Navbar from "./common/components/Navbar";
import {useDispatch} from "react-redux";
import { getAllProductSpecifications} from "./store/app-slice";
import {ThunkDispatch} from "redux-thunk";
import {ProductCategory} from "./__generated__/graphql";

function App() {
    const auth = useAuth();
    const navigate = useNavigate();
    const dispatch: ThunkDispatch<ProductCategory[], any, any> = useDispatch();

    useEffect(() => {
        if (!auth.isAuthenticated() && !auth.isPending()) {
            auth.tryRefreshToken();

            if ( auth.isError() || auth.isEmpty() ) {
                navigate('/auth/login');
            }
        }

        if (auth.isAuthenticated()) {
            if (auth.user?.isRegistering()) {
                navigate('/billing');
            }

            if ( auth.isFullfilled() && auth.user?.isCompleted() ) {
                dispatch( getAllProductSpecifications() );
            }
        }
    }, [auth]);


    return (
        <>
            {auth.isPending() && <Spinner/>}
            <Navbar></Navbar>
            <main className="min-h-screen">
                <Outlet/>
            </main>
            <CustomFooter/>
        </>
    );
}

export default App;
