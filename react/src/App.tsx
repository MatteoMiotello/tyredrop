import './App.css';
import {useEffect} from "react";
import {useDispatch, useSelector} from "react-redux";
import {Outlet, useNavigate} from "react-router-dom";
import {ThunkDispatch} from "redux-thunk";
import {selectUserStatus} from "./modules/auth/store/auth-selector";
import {authRefreshToken} from "./modules/auth/store/auth-slice";

function App() {
    const userStatus = useSelector(selectUserStatus);
    const dispatch: ThunkDispatch<any, any, any> = useDispatch();
    const navigate = useNavigate();

    useEffect(() => {
        if (userStatus.status == null) {
            const refreshToken = window.localStorage.getItem('refresh_token');

            if (!refreshToken) {
                navigate('/auth/login');
                return;
            }

            dispatch(authRefreshToken(refreshToken));
        }

        if ( userStatus.status == 'error' ) {
            navigate( '/auth/login' );
        }

    }, [userStatus]);

    return (
        <>
            <Outlet></Outlet>
        </>
    );
}

export default App;
