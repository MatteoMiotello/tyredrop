import './App.css';
import {useEffect} from "react";
import {useDispatch, useSelector} from "react-redux";
import {Outlet, useNavigate} from "react-router-dom";
import {ThunkDispatch} from "redux-thunk";
import {UserStatus} from "./modules/auth/service/user";
import {selectAuthStatus, selectUser} from "./modules/auth/store/auth-selector";
import {authRefreshToken} from "./modules/auth/store/auth-slice";
import {useAuthenticated} from "./modules/auth/hooks/useAuthenticated";

function App() {
    const authStatus = useSelector(selectAuthStatus);
    const user = useSelector( selectUser );
    const dispatch: ThunkDispatch<any, any, any> = useDispatch();
    const navigate = useNavigate();
    const isAuthenticated = useAuthenticated();

    useEffect(() => {
        if ( authStatus.status == 'error' ) {
            navigate( '/auth/login' );
            return;
        }

        if ( user?.status && user.status == UserStatus.REGISTERING ) {
            navigate( '/auth/billing' );
            return;
        }

        if (!isAuthenticated) {
            const refreshToken = window.localStorage.getItem('refresh_token');

            if (!refreshToken) {
                navigate('/auth/login');
                return;
            }

            dispatch(authRefreshToken(refreshToken));
        }

    }, [authStatus]);

    return (
        <>
            <Outlet></Outlet>
        </>
    );
}

export default App;
