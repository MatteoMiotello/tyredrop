import React from "react";
import Logo from "./Logo";
import Input from "../components-library/Input";
import {useAuth} from "../../modules/auth/hooks/useAuth";
import {FontAwesomeIcon} from "@fortawesome/react-fontawesome";
import {faSearch, faUser} from "@fortawesome/free-solid-svg-icons";
import Button from "../components-library/Button";

const Navbar: React.FC = () => {
    const auth = useAuth();

    return <div className="navbar bg-base-100">
            <Logo width={60}/>
            <div className="input-group">
                <Input className="input-sm m-auto w-96" type="text" name="search" placeholder="Search" />
                <Button>
                    <FontAwesomeIcon icon={faSearch}/>
                </Button>
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
                <ul tabIndex={0} className="mt-3 p-2 shadow menu menu-compact dropdown-content bg-base-100 rounded-box w-52">
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