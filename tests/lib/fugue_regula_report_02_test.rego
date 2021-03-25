# Copyright 2020-2021 Fugue, Inc.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

# This test case produces a report for multiple cloudformation files.
package fugue.regula_report_02_test

import data.fugue.regula
import data.tests.rules.cfn.cloudtrail.inputs.invalid_encryption_infra as input1
import data.tests.rules.cfn.s3.inputs.valid_encryption_infra as input2

mock_input := [
  {
    "filename": "template1.yaml",
    "content": input1.mock_plan_input
  },
  {
    "filename": "template2.yaml",
    "content": input2.mock_plan_input
  }
]

mock_rules := {
  "deny_cloudtrails": {
    "resource_type": "AWS::CloudTrail::Trail",
    "input_type": "cloudformation",
    "allow": false
  },
  "allow_buckets": {
    "resource_type": "AWS::S3::Bucket",
    "input_type": "cloudformation",
    "deny": false
  }
}

report = ret {
  ret := regula.report with input as mock_input with data.rules as mock_rules
}

test_report {
  report.summary.rule_results.PASS == 2
  report.summary.rule_results.FAIL == 1

  report.summary.filenames[_] == "template1.yaml"
  report.summary.filenames[_] == "template2.yaml"

  report.rule_results[i].filename == "template1.yaml"
  report.rule_results[i].resource_type == "AWS::CloudTrail::Trail"
  report.rule_results[i].rule_result == "FAIL"

  report.rule_results[j].filename == "template2.yaml"
  report.rule_results[j].resource_type == "AWS::S3::Bucket"
  report.rule_results[j].rule_result == "PASS"
}
