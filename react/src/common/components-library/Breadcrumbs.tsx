import React, {ReactNode} from "react";
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
        .filter((match) => {
            const handle = match.handle as { crumb: () => ReactNode };

            return Boolean(handle?.crumb);
        })
        .map((match: any) => {
            return match.handle.crumb(match.data);
        });

    return <div className="text-sm breadcrumbs pl-4 bg-base-100 rounded-box m-1">
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