import './App.css';
import {useEffect} from "react";
import {useDispatch} from "react-redux";
import {Outlet, useNavigate} from "react-router-dom";
import {ThunkDispatch} from "redux-thunk";
import Spinner from "./common/components/Spinner";
import {useAuth} from "./modules/auth/hooks/useAuth";
import {UserStatus} from "./modules/auth/service/user";
import {authRefreshToken} from "./modules/auth/store/auth-slice";

function App() {
    const auth = useAuth();
    const dispatch: ThunkDispatch<any, any, any> = useDispatch();
    const navigate = useNavigate();

    useEffect(() => {
        if (!auth.isAuthenticated() && !auth.isPending()) {
            const refreshToken = window.localStorage.getItem('refresh_token');

            if (refreshToken && auth.isEmpty()) {
                dispatch(authRefreshToken(refreshToken));
                return;
            }

            if ( auth.isEmpty() || auth.isError() ) {
                navigate('/auth/login');
            }
        }

        if (auth.isAuthenticated()) {
            if (auth.user?.status && auth.user.status == UserStatus.REGISTERING) {
                navigate('/auth/billing');
            }
        }
    }, [auth]);

    return (
        <>
            {(auth.isEmpty() || auth.isPending()) && <Spinner/>}
            <Outlet></Outlet>
        </>
    );
}

export default App;
