import Navbar from "../components/Navbar";
import NoGroups from "./NoGroupsBody";
import Notes from "./Notes";
import Page from "../components/Page";
import NoGroupsNav from "./NoGroupsNav";

interface BodyProps{
    navbarIsOpen: boolean
    groups: JSX.Element[]
}

function Body(props: BodyProps) {
    let pageContent : JSX.Element;
    let navContent : JSX.Element;

    if (props.groups.length === 0) {
        pageContent = <NoGroups />
        navContent = <NoGroupsNav />
    } else {
        pageContent = <Notes />
        navContent = <NoGroupsNav />
    }

    return (
        <div className="relative lg:flex flex-grow">
            <Navbar isOpen={props.navbarIsOpen}>
                {navContent}
            </Navbar>
            <Page>
                {pageContent}
            </Page>
        </div>
    )
}

export default Body;