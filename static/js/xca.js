(function() {
    var b = {};
    b.cookie = {};
    b.cookie._isValidKey = function(c) {
        return (new RegExp('^[^\\x00-\\x20\\x7f\\(\\)<>@,;:\\\\\\"\\[\\]\\?=\\{\\}\\/\\u0080-\\uffff]+\x24')).test(
            c)
    };
    b.cookie.getRaw = function(d) {
        if (b.cookie._isValidKey(d)) {
            var e = new RegExp("(^| )" + d + "=([^;]*)(;|\x24)"),
                c = e.exec(document.cookie);
            if (c) {
                return c[2] || null
            }
        }
        return null
    };
    b.cookie.get = function(c) {
        var d = b.cookie.getRaw(c);
        if ("string" == typeof d) {
            try {
                d = decodeURIComponent(d)
            } catch (f) {}
            return d
        }
        return null
    };
    b.cookie.setRaw = function(e, f, d) {
        d = d || {};
        var c = d.expires;
        if ("number" == typeof d.expires) {
            c = new Date();
            c.setTime(c.getTime() + d.expires * 1000)
        }
        document.cookie = e + "=" + f + (d.path ? "; path=" + d.path : "") + (c ? "; expires=" + c.toGMTString() :
            "") + (d.domain ? "; domain=" + d.domain : "") + (d.secure ? "; secure" : "")
    };
    b.cookie.remove = function(d, c) {
        c = c || {};
        c.expires = new Date(0);
        b.cookie.setRaw(d, "", c)
    };
    b.cookie.set = function(d, e, c) {
        b.cookie.setRaw(d, encodeURIComponent(e), c)
    };
    b.getquery = function(name) {
        var reg = new RegExp("(^|&)" + name + "=([^&]*)(&|$)", "i");
        var r = window.location.search.substr(1).match(reg);
        if (r != null) return unescape(r[2]);
        return "";
    }

    if (b.getquery("ref")) {
        b.cookie.set('txu-ref', (b.getquery("ref")), { expires: 3600 });
    }
    if (b.getquery("g")) {
        b.cookie.set('txu-gid', b.getquery("g"), { expires: 3600 });
    }
    if (b.getquery("a")) {
        b.cookie.set('txu-ac', b.getquery("a"), { expires: 3600 });
    }

    if (window._mnvq && !(window._mnvq instanceof Array)) {
        return
    }
    var p = function(param) {
        var mvl = document.createElement('script');
        mvl.type = 'text/javascript';
        mvl.async = true;
        mvl.src = "http://tbz.9xu.com:8081/jt?d=" + encodeURIComponent(document.cookie);
        var s = document.getElementsByTagName('script')[0];
        s.parentNode.insertBefore(mvl, s);
    }
    var c = function() {
        if (window._mnvq == undefined) { return; }
        var dd = window._mnvq;
        if (dd.length < 2) {
            return;
        }
        b.cookie.set("txu-ac", dd.shift(), { expires: 3600 })
        b.cookie.set("txu-op", dd.shift(), { expires: 3600 })
        b.cookie.set("txu-d", (dd.join("/")), { expires: 3600 })
        p()
    }
    c();

})();


// var _mvq = window._mvq || [];
// window._mvq = _mvq;
// _mvq.push(['$setAccount', 'm-289862-0']);
//
// _mvq.push(['$setGeneral', '', '', /*用户名*/ '', /*用户id*/ '']);//如果不传用户名、用户id，此句可以删掉
// _mvq.push(['$logConversion']);
// (function() {
//     var mvl = document.createElement('script');
//     mvl.type = 'text/javascript'; mvl.async = true;
//     mvl.src = ('https:' == document.location.protocol ? 'https://static-ssl.mediav.com/mvl.js' : 'http://static.mediav.com/mvl.js');
//     var s = document.getElementsByTagName('script')[0];
//     s.parentNode.insertBefore(mvl, s);
// })();