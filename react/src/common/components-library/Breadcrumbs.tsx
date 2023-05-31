import React from "react";
import {useMatches} from "react-router-dom";

type BreadcrumbLink = {
    title: string
    link: string | undefined
}

type BreadcrumbsProps = {
    links: BreadcrumbLink[]
}

const Breadcrumbs: React.FC = (props) => {
    const matches = useMatches();
    const crumbs = matches
        // first get rid of any matches that don't have handle and crumb
        .filter((match) => Boolean(match.handle?.crumb))
        // now map them into an array of elements, passing the loader
        // data to each one
        .map((match) => match.handle.crumb(match.data));

    return <div className="text-sm breadcrumbs pl-4 bg-base-200">
        <ul>
            {
                crumbs.map((link, key) => {
                    return <li key={key}>{link}</li>;
                })
            }
        </ul>
    </div>;
};

export default Breadcrumbs;