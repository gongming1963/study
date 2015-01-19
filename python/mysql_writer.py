import MySQLdb

class MysqlWriter:
    def __init__(self):
        self.connect = None

    def _connect(self):
        if self.connect:
            return
        self.connect = MySQLdb.connect(**meta)
        self.connect.autocommit(0)

    def write_raw(self, sql):
        self._connect()
        cursor = self.connect.cursor()
        cursor.execute(sql)
        self.connect.commit()
        cursor.close()

mysql_writer = MysqlWriter()
