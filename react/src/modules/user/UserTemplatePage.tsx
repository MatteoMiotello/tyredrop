import { faUser} from "@fortawesome/free-solid-svg-icons";
import {FontAwesomeIcon} from "@fortawesome/react-fontawesome";
import React from "react";
import { Outlet} from "react-router-dom";
import UserMenu from "./components/UserMenu";

const UserTemplatePage: React.FC = () => {
    return <div className="">
        <div className="flex flex-col md:flex-row min-h-screen items-stretch">
            <div className="p-4 flex flex-col bg-base-100 rounded-box m-1">
                <div className="avatar placeholder aspect-square p-1 mx-auto border-2">
                    <div className="bg-neutral-focus text-neutral-content rounded-full ">
                        <span className="text-6xl"> <FontAwesomeIcon className="w-36" icon={faUser}/> </span>
                    </div>
                </div>
                <UserMenu/>
            </div>
            <div className="w-full p-1">
                <Outlet/>
            </div>
        </div>
    </div>;
};

export default UserTemplatePage;