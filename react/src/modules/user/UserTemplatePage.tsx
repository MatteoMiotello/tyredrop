import { faUser} from "@fortawesome/free-solid-svg-icons";
import {FontAwesomeIcon} from "@fortawesome/react-fontawesome";
import React from "react";
import { Outlet} from "react-router-dom";
import UserMenu from "./components/UserMenu";

const UserTemplatePage: React.FC = () => {
    return <div className="md:p-24 p-4">
        <div className="grid grid-cols-4 border-2 rounded-box">
            <div className="h-full bg-base-300 p-4 flex flex-col">
                <div className="avatar placeholder h-full p-1 mx-auto">
                    <div className="bg-neutral-focus text-neutral-content rounded-full ">
                        <span className="text-6xl"> <FontAwesomeIcon className="w-36" icon={faUser}/> </span>
                    </div>
                </div>
                <UserMenu/>
            </div>
            <div className="col-span-3 p-4">
                <Outlet/>
            </div>
        </div>
    </div>;
};

export default UserTemplatePage;