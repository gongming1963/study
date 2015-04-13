#!/usr/local/bin/python
#coding=utf-8
import sys
reload(sys)
sys.setdefaultencoding('utf-8')
import traceback
from optparse import OptionParser

#PARAMS = {'ori_table':'detail.compdeal_daily',
#        'new_table':'ba_cis.compdeal_daily_todetect',
#        'primary_key':['date', 'dealid'],
#        'keys':['addtime','volume','status','ord','origin' ],
#        'diff_key':['']}
#
PARAMS = {'ori_table':'test.compdeal_daily',
        'new_table':'summary.compdeal_daily',
        'primary_key':['date', 'siteid', 'cityid'],
        'keys':['date','siteid','websiteid','cityid','dealcount','coupondealcount','deliverydealcount','bigdealcount','bigdealdealcount','showdays','couponshowdays','deliveryshowdays','bigdealshowdays','quantity','couponquantity','deliveryquantity','bigdealquantity','revenue','couponrevenue','deliveryrevenue','bigdealrevenue','revenue_origin','couponrevenue_origin','deliveryrevenue_origin','bigdealrevenue_origin','totalvalue','totalcut','bigdealtotalvalue','bigdealtotalcut','coupontotalvalue','coupontotalcut','deliverytotalvalue','deliverytotalcut'],
        'diff_key':['dealcount']}

def generate_sample_sql():
    """
    generate sql to get diff sample
    """
    sample_sql = "select "
    key_item = []
    table_list = ['ori', 'new']
    KEY_PATTERN = "%(table_name)s.%(key_name)s %(table_name)s_%(key_name)s"
    for key in PARAMS['keys']:
        for table in table_list:
            key_item.append(KEY_PATTERN % {'table_name':table, 'key_name':key})
    sample_sql += ',\n'.join(key_item)
    sample_sql += " \nfrom %(ori_table)s ori FULL OUTER JOIN %(new_table)s new \nON " % PARAMS

    on_item = []
    ON_PATTERN = "ori.%(primary_key)s = new.%(primary_key)s"
    for primary_key in PARAMS['primary_key']:
        on_item.append(ON_PATTERN % {'primary_key':primary_key})
    sample_sql += ' and '.join(on_item)

    where_item = []
    WHERE_PATTERN = "ori.%(diff_key)s != new.%(diff_key)s"
    for diff_key in PARAMS['diff_key']:
        on_item.append(WHERE_PATTERN % {'diff_key':diff_key})
    sample_sql += " \nwhere " + ' and '.join(on_item) + " \nlimit 1000;"

    print sample_sql

def generate_diff_sql():
    """
    generate diff sql
    """
    TABLE_NUM = "sum(if(%(table_name)s.%(primary_key)s is not NULL, 1, 0)) %(table_name)s_num,\n"
    diff_sql = "select count(*) total_num,\n"
    diff_sql += TABLE_NUM % {'primary_key':PARAMS['primary_key'][0], 'table_name':'ori'}
    diff_sql += TABLE_NUM % {'primary_key':PARAMS['primary_key'][0], 'table_name':'new'}

    key_item = []
    KEY_DIFF = "sum(if(ori.%(key)s= new.%(key)s,1,0))  %(key)s_num"
    for key in PARAMS['keys']:
        key_item.append(KEY_DIFF % {'key':key})
    diff_sql += ',\n'.join(key_item)

    diff_sql += " \nfrom %(ori_table)s ori FULL OUTER JOIN %(new_table)s new \nON " % PARAMS

    on_item = []
    ON_PATTERN = "ori.%(primary_key)s = new.%(primary_key)s"
    for primary_key in PARAMS['primary_key']:
        on_item.append(ON_PATTERN % {'primary_key':primary_key})
    diff_sql += ' and '.join(on_item)
    diff_sql += ";"

    print diff_sql

def main():
    """
    main
    """
    options = set_usage()
    if options.function == 'sample':
        generate_sample_sql()
    else:
        generate_diff_sql()

def set_usage() :
    """
    usage 
    """
    parser = OptionParser(usage='usage: ')
    parser.add_option("-f", dest="function", help="")
    params, args = parser.parse_args()
    if params.function not in set(['diff', 'sample']):
        print "error function"
        exit(-1)

    return  params

if __name__ == "__main__":
    main()

