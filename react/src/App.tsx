import './App.css';
import {useEffect} from "react";
import {useDispatch} from "react-redux";
import {Outlet, useNavigate} from "react-router-dom";
import {ThunkDispatch} from "redux-thunk";
import CustomFooter from "./common/components/CustomFooter";
import Spinner from "./common/components/Spinner";
import {useAuth} from "./modules/auth/hooks/useAuth";
import {UserStatus} from "./modules/auth/service/user";

function App() {
    const auth = useAuth();
    const dispatch: ThunkDispatch<any, any, any> = useDispatch();
    const navigate = useNavigate();

    useEffect(() => {
        if (!auth.isAuthenticated() && !auth.isPending()) {
            auth.tryRefreshToken();

            if ( auth.isError() ) {
                navigate('/auth/login');
            }
        }

        if (auth.isAuthenticated()) {
            if (auth.user?.status && auth.user.status == UserStatus.REGISTERING) {
                navigate('/billing');
            }
        }
    }, [auth]);

    return (
        <>
            {auth.isPending() && <Spinner/>}
            <main className="min-h-screen lg:p-24 p-4 h-full w-full flex flex-col">
                <Outlet/>
            </main>
            <CustomFooter/>
        </>
    );
}

export default App;
