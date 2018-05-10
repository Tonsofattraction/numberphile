from contextlib import contextmanager
import time

@contextmanager
def log_time_context(tag='', longer=0):
    """

    :param tag: log time preceeded with this tag
    :param longer: log only if execution took longer than this (seconds)
    :return:
    """
    start = time.time()
    yield
    end = time.time()
    if end - start > longer:
        logging.error('CONTEXT TIMER %s: %s' % (tag, round(end - start, 3)))