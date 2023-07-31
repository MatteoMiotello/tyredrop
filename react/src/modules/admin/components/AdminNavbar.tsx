import { faRightFromBracket, faUser} from "@fortawesome/free-solid-svg-icons";
import {FontAwesomeIcon} from "@fortawesome/react-fontawesome";
import React from "react";
import {useDispatch} from "react-redux";
import {Link} from "react-router-dom";
import Menu from "../../../common/components-library/Menu";
import MainLogo from "../../../common/components/Logo";
import {logout} from "../../auth/store/auth-slice";

const AdminNavbar: React.FC = () => {
    const dispatch  = useDispatch();

    return <div className="navbar bg-primary rounded-box">
        <MainLogo width={100}/>
        <div className="text-4xl text-white uppercase font-semibold ml-4">
            <h1>
                Area privata
            </h1>
        </div>
        <div className="ml-auto">
            <div className="flex-none gap-2">
                <div className="dropdown dropdown-end z-50">
                    <Menu.Dropdown label={<div className="avatar placeholder h-full p-1">
                        <div className="bg-neutral-focus text-neutral-content rounded-full">
                            <span className="text-xl"> <FontAwesomeIcon icon={faUser}/> </span>
                        </div>
                    </div>}>
                        <Menu.Item>
                            <Link to="/user">
                                <FontAwesomeIcon icon={faUser}/> Principale
                            </Link>
                        </Menu.Item>
                        <Menu.Item>
                            <a onClick={() => dispatch(logout())}>
                                <FontAwesomeIcon icon={faRightFromBracket}/> Esci
                            </a>
                        </Menu.Item>
                    </Menu.Dropdown>
                </div>
            </div>
        </div>
    </div>;
};

export default AdminNavbar;