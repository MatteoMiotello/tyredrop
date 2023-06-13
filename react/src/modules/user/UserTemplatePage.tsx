import { faUser} from "@fortawesome/free-solid-svg-icons";
import {FontAwesomeIcon} from "@fortawesome/react-fontawesome";
import React from "react";
import { Outlet} from "react-router-dom";
import UserMenu from "./components/UserMenu";

const UserTemplatePage: React.FC = () => {
    return <div className="">
        <div className="grid grid-cols-12 rounded-box min-h-screen bg-base-200">
            <div className="col-span-2 h-full p-4 flex flex-col">
                <div className="avatar placeholder aspect-square p-1 mx-auto border-2">
                    <div className="bg-neutral-focus text-neutral-content rounded-full ">
                        <span className="text-6xl"> <FontAwesomeIcon className="w-36" icon={faUser}/> </span>
                    </div>
                </div>
                <UserMenu/>
            </div>
            <div className="col-span-10 p-4 bg-base-100 rounded-tl-box">
                <Outlet/>
            </div>
        </div>
    </div>;
};

export default UserTemplatePage;