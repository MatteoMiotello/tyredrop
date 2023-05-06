import React from "react";
import Field from "../components-library/Input";
import MainLogo from "./Logo";
import {FontAwesomeIcon} from "@fortawesome/react-fontawesome";
import {faSearch, faUser} from "@fortawesome/free-solid-svg-icons";
import {useAuth} from "../../modules/auth/hooks/useAuth";

const Navbar: React.FC = () => {
    const auth = useAuth();

    return <div className="navbar bg-base-100">
        <MainLogo width={60}/>
        <div className="input-group">
            <Field.FormControl className="m-auto px-4 md:w-1/2 w-full">
                <Field.InputGroup size="sm">
                    <Field.Input type="text" name="search" placeholder="Ricerca facile" size="sm"/>
                    <span>
                        <FontAwesomeIcon icon={faSearch}/>
                    </span>
                </Field.InputGroup>
            </Field.FormControl>
        </div>
        <div className="flex-none gap-2">
            <div className="dropdown dropdown-end">
                <label tabIndex={0} className="btn btn-ghost btn-circle avatar">
                    <div className="avatar placeholder h-full p-1">
                        <div className="bg-neutral-focus text-neutral-content rounded-full">
                            <span className="text-xl"> <FontAwesomeIcon icon={faUser}/> </span>
                        </div>
                    </div>
                </label>
                <ul tabIndex={0}
                    className="mt-3 p-2 shadow menu menu-compact dropdown-content bg-base-100 rounded-box w-52">
                    {
                        auth && <li>
                        { auth.user?.getCompleteName() }
                    </li>
                    }
                    <li>
                        <a className="justify-between">
                            Profile
                            <span className="badge">New</span>
                        </a>
                    </li>
                    <li><a>Settings</a></li>
                    <li><a>Logout</a></li>
                </ul>
            </div>
        </div>
    </div>;
};

export default Navbar;