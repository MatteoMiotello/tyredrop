import './App.css';
import {useEffect} from "react";
import {useTranslation} from "react-i18next";
import {Outlet, useNavigate} from "react-router-dom";
import Breadcrumbs from "./common/components-library/Breadcrumbs";
import CustomFooter from "./common/components/CustomFooter";
import Spinner from "./common/components/Spinner";
import {useToast} from "./hooks/useToast";
import {useAuth} from "./modules/auth/hooks/useAuth";
import Navbar from "./common/components/Navbar";
import {useDispatch} from "react-redux";
import {fetchUserAddresses} from "./modules/user/store/user-slice";
import {getAllProductSpecifications} from "./store/app-slice";
import {ThunkDispatch} from "redux-thunk";
import {ProductCategory} from "./__generated__/graphql";
import {fetchCartItems} from "./modules/cart/store/cart-slice";

function App() {
    const auth = useAuth();
    const navigate = useNavigate();
    const {t} = useTranslation();
    const {setError} = useToast();
    const dispatch: ThunkDispatch<ProductCategory[], any, any> = useDispatch();

    useEffect(() => {
        if (!auth.isAuthenticated() && !auth.isPending()) {
            auth.tryRefreshToken();

            if ( auth.isError() ) {
                navigate('/auth/login');
                setError( t( 'login.error_redirect' ) );
                return;
            }

            if (auth.isEmpty()) {
                navigate('/auth/login');
                return;
            }
        }

        if (auth.isAuthenticated()) {
            if (auth.user?.isNotConfirmed()) {
                navigate('/not_confirmed');
            }

            if (auth.user?.isRegistering()) {
                navigate('/billing');
            }

            if (auth.isFullfilled() && auth.user?.isCompleted()) {
                dispatch(getAllProductSpecifications());
                dispatch(fetchCartItems());
                dispatch(fetchUserAddresses());
            }
        }
    }, [auth]);

    return (
        <>
            {auth.isPending() && <Spinner/>}
            <Navbar></Navbar>
            <Breadcrumbs></Breadcrumbs>
            <main className="min-h-screen bg-base-100">
                <Outlet/>
            </main>
            <CustomFooter/>
        </>
    );
}

export default App;
