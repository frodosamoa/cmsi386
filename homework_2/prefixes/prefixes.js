/**
 * This function passes a string and a prefix index as paramenters.
 * It then writes successive prefixes of the value entered into the
 * input field as a <div> element/
 */

function prefixes(s, i) {
    if (i <= s.length) {
        $("body").append("<div>" + (i === 0 ? "&nbsp;" : s.substring(0,i)) 
        	+ "</div>");
        setTimeout(function () { prefixes(s, i+1) }, 1000);
    }
}