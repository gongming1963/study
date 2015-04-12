'''
usage : Class of download page according to url.
author : gongming
date : 2013-12-30
'''

import time
import httplib
import logging
import cPickle
import string
import traceback

from urlparse import urlparse, urljoin
from httplib import HTTPConnection

class Grabber():
    '''
    Grabber class.
    '''
    def __init__(self):
        self._agent = 'Mozilla/5.0 (Windows; U; Windows NT 5.1; zh-CN; rv:1.9.1.5) Gecko/20091102 Firefox/3.5.5'
        self._accept = 'text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8'
        self._lang = 'zh-cn,zh;q=0.5'
        self._time_out = 30

    def get_ip(self, hostname):
        '''
        Resolve hostname.
        '''
        ip_addr = ""
        try:
            ip_addr = httplib.socket.gethostbyname(hostname)
        except Exception, ex:
            logging.warn("get ip of %s exception: %s" % (hostname, ex.message))
            ip_addr = ""
        
        return ip_addr        
        
    def grab(self, url):
        '''
        Grab page by url, ip.
        '''
        if not url :
            logging.warn("url is None!")
            return None

        print "*********************************************START*********************************************"
        print "Get page from url:%s" % url

        port = 80
        urltuple = urlparse(url)
        if urltuple.port:
            port = urltuple.port

        host_ip = self.get_ip(urltuple.netloc)

        content = ""
        headers = ""

        conn = None
        try:
            httplib.socket.setdefaulttimeout(self._time_out)
            conn = HTTPConnection(host_ip, port)
        except Exception, ex:
            logging.warn('url: %s, connect to %s:%d exception: %s' \
                    % (url, host_ip, port, ex.message))
            return content

        try:
            absurl = urltuple.path + "?" + urltuple.query
            conn.putrequest('GET', absurl, 1, 1)
            conn.putheader("HOST", urltuple.hostname)
            conn.putheader("Accept", self._accept)
            conn.putheader("Accept-Language", self._lang)
            conn.putheader("Accept-Charset", "utf-8")
            conn.putheader("User-Agent", self._agent)
            conn.endheaders()
            conn.send('')
            response = conn.getresponse()

            print "\nHEADER:"
            for key in response.getheaders():
                print "%s:%s" % (key[0],key[1])

            print self.get_content(response)
        except Exception, ex:
            logging.warn('get page of {%s} exception: %s' % (url, ex.message))
            time.sleep(10)
        finally:
            conn.close()
        
        return content

    def get_content(self, response):
        """
        get content from response
        """
        print "\nCONTENT:"
        res_status = response.status
        if res_status == 302:
            print response.read()
            content = self.grab(response.getheader('location'))
        else:
            content = response.read()
            if response.chunked:
                content = decode_chunked(content)
        return content

def decode_chunked(data):
    offset = 0
    encdata = ''
    newdata = ''
    offset = string.index(data, "\r\n\r\n") + 4 # get the offset 
    encdata =data[offset:]
    try:
        while (encdata != ''):
            off = int(encdata[:string.index(encdata,"\r\n")],16)
            if off == 0:
                break
            encdata = encdata[string.index(encdata,"\r\n") + 2:]
            newdata = "%s%s" % (newdata, encdata[:off])
            encdata = encdata[off+2:]
                             
    except:
       line = traceback.format_exc()
       logging.warn("Exception! %s" %line) # probably indexes are wrong
    return newdata

if __name__ == "__main__":
    GRABBER = Grabber()
    URL = "http://t.cn/h5mwx"
    print GRABBER.grab(URL)
