import { FontAwesomeIcon } from "@fortawesome/react-fontawesome";
import { faExclamationCircle } from "@fortawesome/free-solid-svg-icons";

function NoGroupsNav() {
    return (
        <div>
            <FontAwesomeIcon icon={faExclamationCircle} size="10x" className="block mt-20 mb-5 mx-auto opacity-10"/>
            <div className="text-center opacity-20 mx-3">
                Your groups will appear here once you have joined one.
            </div>
        </div>
    )
}

export default NoGroupsNav;