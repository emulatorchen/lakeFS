"""
    lakeFS API

    lakeFS HTTP API  # noqa: E501

    The version of the OpenAPI document: 1.0.0
    Contact: services@treeverse.io
    Generated by: https://openapi-generator.tech
"""


import sys
import unittest

import lakefs_client
from lakefs_client.model.pagination import Pagination
from lakefs_client.model.pull_request import PullRequest
globals()['Pagination'] = Pagination
globals()['PullRequest'] = PullRequest
from lakefs_client.model.pull_requests_list import PullRequestsList


class TestPullRequestsList(unittest.TestCase):
    """PullRequestsList unit test stubs"""

    def setUp(self):
        pass

    def tearDown(self):
        pass

    def testPullRequestsList(self):
        """Test PullRequestsList"""
        # FIXME: construct object with mandatory attributes with example values
        # model = PullRequestsList()  # noqa: E501
        pass


if __name__ == '__main__':
    unittest.main()
