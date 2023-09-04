import React from "react";
import CustomFooter from "./common/components/CustomFooter";
import Logo from "./common/components/Logo";
import {Link, Outlet} from "react-router-dom";
import {useAuth} from "./modules/auth/hooks/useAuth";
import {logout} from "./modules/auth/store/auth-slice";
import {FontAwesomeIcon} from "@fortawesome/react-fontawesome";
import {faRightFromBracket} from "@fortawesome/free-solid-svg-icons";
import {useDispatch} from "react-redux";
import {ThunkDispatch} from "redux-thunk";
import {Button} from "./common/components/shelly-ui";

const CommonTemplate: React.FC = () => {
    const auth = useAuth();
    const dispatch = useDispatch<ThunkDispatch<any, any, any>>();

    return <div>
        <div className="navbar w-full">
            <Logo className="navbar-start w-24"/>
            <div className="ml-auto flex gap-2">
                {
                    auth.isLoggedIn() ? <Link to="/" className="btn btn-primary"> Home </Link> : <Link to="/auth/login" className="btn btn-primary"> Login </Link>
                }
                {
                    auth.isLoggedIn() && <Button buttonType="ghost" outline onClick={() => dispatch(logout())}>
                        <FontAwesomeIcon icon={faRightFromBracket}/> Esci
                    </Button>
                }
            </div>
        </div>
        <div className="p-4 md:p-24 md:px-56">
            <Outlet/>
        </div>
        <CustomFooter/>
    </div>;
};

export default CommonTemplate;