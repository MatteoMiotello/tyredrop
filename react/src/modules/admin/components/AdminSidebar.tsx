import {faBoxesStacked, faEuro, faHome, faUsers} from "@fortawesome/free-solid-svg-icons";
import {FontAwesomeIcon} from "@fortawesome/react-fontawesome";
import React from "react";
import {Link} from "react-router-dom";
import {twMerge} from "tailwind-merge";
import Menu from "../../../common/components-library/Menu";

type AdminSidebarProps = {
    className?: string
}

const AdminSidebar: React.FC<AdminSidebarProps> = ( {className} ) => {
    const classNames = twMerge(
        "bg-base-100 rounded-box shadow",
        
        className
    );

    return <div className={classNames}>
        <Menu>

            <Menu.Item>
                <Link to={'/admin'}> <FontAwesomeIcon icon={faHome}/> Home </Link>
            </Menu.Item>
            <Menu.Item>
                <Link to={'/admin/price'}> <FontAwesomeIcon icon={faEuro}/> Prezzi </Link>
            </Menu.Item>
            <Menu.Item>
                <Link to={'/admin/user'}> <FontAwesomeIcon icon={faUsers}/> Utenti </Link>
            </Menu.Item>
            <Menu.Item>
                <Link to={'/admin/order'}> <FontAwesomeIcon icon={faBoxesStacked}/> Ordini </Link>
            </Menu.Item>
        </Menu>
    </div>;
};

export default AdminSidebar;