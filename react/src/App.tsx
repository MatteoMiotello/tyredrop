import {useEffect} from "react";
import {useTranslation} from "react-i18next";
import {Outlet, useNavigate} from "react-router-dom";
import Breadcrumbs from "./common/components-library/Breadcrumbs";
import CustomFooter from "./common/components/CustomFooter";
import {useAuth} from "./modules/auth/hooks/useAuth";
import Navbar from "./common/components/Navbar";
import {useDispatch} from "react-redux";
import {authRefreshToken} from "./modules/auth/store/auth-slice";
import {fetchUserAddresses} from "./modules/user/store/user-slice";
import {getAllProductSpecifications} from "./store/app-slice";
import {ThunkDispatch} from "redux-thunk";
import {ProductCategory} from "./__generated__/graphql";
import {fetchCartItems} from "./modules/cart/store/cart-slice";

function App() {
    const auth = useAuth();
    const navigate = useNavigate();
    const {t} = useTranslation();
    const dispatch: ThunkDispatch<ProductCategory[], any, any> = useDispatch();

    useEffect(() => {
        if (auth.unknownStatus()) {
            const refreshToken = window.localStorage.getItem('refresh_token');

            if (refreshToken) {
                dispatch(authRefreshToken(refreshToken));
                return;
            }

            navigate('/auth/login');
            return;
        }

        if (auth.isNotLoggedIn()) {
            navigate('/auth/login');
            return;
        }

        if (auth.isLoggedIn()) {
            if (auth.isAdmin()) {
                navigate('/admin');
                return;
            }

            if (auth.user?.isNotConfirmed()) {
                navigate('/not_confirmed');
                return;
            }

            if (auth.user?.isRegistering()) {
                navigate('/billing');
                return;
            }

            dispatch(getAllProductSpecifications());
            dispatch(fetchCartItems());
            dispatch(fetchUserAddresses());
        }
    }, [auth]);

    return (
        <>
            <div className="xl:px-56">
                <div className="p-1">
                    <Navbar></Navbar>
                </div>
                <Breadcrumbs></Breadcrumbs>
                <div className="min-h-screen">
                    <Outlet/>
                </div>
            </div>
            <CustomFooter/>
        </>
    );
}

export default App;
