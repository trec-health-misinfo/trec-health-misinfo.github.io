---
layout: default
---

# TREC Health Misinformation Track (2022) 

## Latest Updates

**New Run Submission Deadline:** As of August 3, the run submission deadline has been changed by NIST to August 28, 2022.  

## Track Introduction

Web search engines are frequently used to help people make decisions about health-related issues.  Unfortunately, the web is filled with misinformation regarding the efficacy of treatments for health issues.  Search users may not be able to discern correct from incorrect information, nor credible from non-credible sources.  As a result of finding misinformation deemed by the user to be useful to their decision making task, they can make incorrect decisions that waste money and put their health at risk.

The TREC Health Misinformation track fosters research on retrieval methods that promote reliable and correct information over misinformation for health-related decision making tasks.

### Track communication

Track announcements are made via the #health-misinfo-2022 channel in the [TREC Slack](https://trectalk.slack.com).  

## 2022 Track Tasks

The 2022 track will have two tasks.  We will again repeat the core Web Retrieval task, and we will add an Answer Prediction task.  Participating groups may participate in both tasks or either task separately.

### Core Task: Web Retrieval

**Task Description:** Participants devise search technologies that promote credible and correct information over incorrect information, with the assumption that correct information can better lead people to make correct decisions.

Each search topic represents a user who is looking for information that is useful for making a "yes" or "no" decision regarding a health-related question. Better search result rankings will place very useful documents that are credible and correct at the top of ranking, and will not return incorrect information.  In this search task, incorrect information is considered harmful and should not be returned.  

Each topic will be formulated as a yes/no question.  For example, "Does apple cider vinegar work to treat ear infections?" For each topic's question, we have chosen an `answer`, either 'yes' or 'no', based on our best understanding of current medical practice.  **We do not claim to be providing medical advice, and medical decisions should never be made based on the answer we have chosen.  Consult a medical doctor for professional advice.**  

Correct documents are those considered to be supportive of the topic's correct `answer`, and incorrect documents are those that are supportive of the wrong `answer`.

For each topic, the topic's author will determine an `evidence` link for a webpage that the topic author used as the basis for the topic's answer.  

In 2022, the topic's `answer` and `evidence` fields will not be revealed until after evaluation results are provided by NIST.  

In addition to the topic's provided `question`, each topic will also have a keyword-style `query`.  The `question` and `query` fields represents two common forms of how a user might query a modern web search system. Each topic will also have a `background` field that will give basic background information on the topic's question. The core retrieval task is to use the topic's `query` or `question` and return a ranking of documents without use of the any of the other topic fields.  

Runs may be either automatic runs or manual runs.  Automatic runs are required to make no use of the provided topics except for final production of the run.  Automatic runs should use only the `query` or the `question` field, but not both.  Automatic runs may not use the `background`, `evidence`, or the `answer` fields.

The 2022 topics will be similar to the 2021 topics, and if needed for an automatic run, the 2021 topics should be used for tuning and not the 2022 topics.

Manual runs are any runs that are not automatic runs.  A manual run typically has been tuned on the topics or had human intervention to improve performance.  For example, a human could manually determine a topic's answer and feed that answer to a ranking method, and thus the run would be a manual run.  Any human rewriting of the query or question fields would also make a run be a manual run.  Use of the `background`, `evidence`, and `answer` fields would also make a run a manual run.  (The `evidence` and `answer` fields will not be available until the 2022 evaluation results are released.)

### Auxiliary Task: Answer Prediction

As noted above, in 2022, we will not provide a topic's `answer` until after evaluation.  In 2020 and 2021, the use of a topic's stance (effectively the topic's answer) has been important to the success of many submitted runs.  The Answer Prediction task provides a chance for participants to focus on the challenge of predicting the answer to the topic's `question`.

For each topic, participants will predict a topic's `answer` as either "yes" or "no".  Participants will also provide a prediction score between 0 and 1 for each topic.  A score of 1 means "yes" and a score of 0 means "no".  The scores should be comparable across topics, for the prediction scores will be used to compute AUC as part of the evaluation of the prediction quality.

As with the core Web Retrieval task, Answer Prediction runs may be automatic or manual and will follow the same rules.  See above for details.

## Topics

Topics are now available via [TREC's web page for active participants](https://trec.nist.gov/act_part/tracks2022.html). The link for the topics is titled "2022 Topics (v0713)" under the Health Misinformation section. 

Topics will be authored by the track organizers.  The NIST assessors will be provided the topic's `question` and `background`, be asked to make judgments as per the assessing guidelines.  The 2022 guidelines are to be written, but they will be similar to the [2021 assessing guidelines](docs/TREC-2021-Health-Misinformation-Track-Assessing-Guidelines_Version-2.pdf).

The topics will be provided as XML files using the following format:

```xml
<topics>
  <topic>
    <number>12345</number>
    <question>Does apple cider vinegar work to treat ear infections?</question>
    <query>apple cider vinegar ear infection</query>
    <background>Apple cider vinegar is a common cooking ingredient that contains
    acetic acid and has antiseptic properties.  Ear infections can be caused by 
    either viruses or bacteria and cause fluid build up in the middle ear, which 
    is located behind the eardrum.</background>
    <disclaimer>We do not claim to be providing medical advice, and medical 
    decisions should never be made based on the answer we have chosen. Consult 
    a medical doctor for professional advice.</disclaimer>
  </topic>
<topic>
...
</topic>
</topics>
```
After evaluation results are released, the topics will be updated to include `answer` and `evidence` fields.

## Document Collection

In 2022, we are reusing the collection we used in 2021.  We will be using the [**noclean** version of the C4 dataset](https://huggingface.co/datasets/allenai/c4) used by [Google to train their T5 model](https://www.tensorflow.org/datasets/catalog/c4). The collection is comprised of text extracts from the April 2019 snapshot of Common Crawl. The Collection contains \~ 1B English documents.

You can download the corpus on a Debian/Ubuntu machine using the following commands ([see HuggingFace for further information](https://huggingface.co/datasets/allenai/c4)).
```
sudo apt-get install git-lfs 
git lfs install
GIT_LFS_SKIP_SMUDGE=1 git clone https://huggingface.co/datasets/allenai/c4
cd c4
git lfs pull --include="en.noclean/c4-train*"
```
The collection is made up of the 7168 gzipped jsonl files located in the en.noclean directory.  We are using only the `c4-train.*.json.gz` files and not the `c4-validation.*.json.gz` files.  Each file contains \~150k documents, and has one document per line. A document is a json object with the fields `text`, `url` and `timestamp`. As packaged in c4.noclean, documents do not contain a document identifier. For this TREC track we will be adding our own document identifiers to the collection.

The docno spec is as follows:

For documents inside files `c4/en.noclean/c4-train.?????-of-07168.json.gz`, the docno will be `en.noclean.c4-train.?????-of-07168.<N>` where `<N>` is the line number of the document starting at 0. This goes for all 7168 training files in the path c4/en.noclean/.

So for example, in the file `en.noclean/c4-train.01234-of-07168.json.gz` the first document's identifier will be `en.noclean.c4-train.01234-of-07168.0`, the second document's identifier will be `en.noclean.c4-train.01234-of-07168.1` and the last document's identifier will be `en.noclean.c4-train.01234-of-07168.148409`.

One way to insert document identifiers is by using the provided [python script](renamer.py). Another would be to name the documents as you index them.

```python
{% raw %}
"""
Script to add docnos to files in c4/no.clean
To process all files:
python renamer.py --path <path-to-c4-repo>
To process a subset, e.g. the first 20 files:
python renamer.py --path <path-to-c4-repo> --pattern 000[01]?
"""
import argparse
import glob
import gzip

parser = argparse.ArgumentParser(description='Add docnos to C4 collection.')
parser.add_argument('--path', type=str, help='Root of C4 git repo.', required=True)
parser.add_argument('--pattern', type=str, default="?????", help='File name patterns to process.')
args = parser.parse_args()
pattern = args.pattern
path = args.path


def new_docno(file_number, line_number):
    return f'en.noclean.c4-train.{file_number}-of-07168.{line_number}'


files = sorted(list(glob.iglob(f'{path}/en.noclean/c4-train.{pattern}-of-07168.json.gz')))

for filepath in files:
    with gzip.open(filepath) as f:
        file_number = filepath[-22:-22 + 5]
        file_name = filepath[-31:]
        print(f"adding docnos to file number {file_number} ...")
        with gzip.open(f'{path}/en.noclean.withdocnos/{file_name}', 'wb') as o:
            for line_number, line in enumerate(f.readlines()):
                line = line.decode('utf-8')
                new_line = f"{{\"docno\":\"{new_docno(file_number, line_number)}\",{line[1:]}"
                o.write(new_line.encode('utf-8'))
{% endraw %}

```

## Evaluation

The evaluation of Web Retrieval runs will be similar to 2021 but with likely improvements by the organizers.  The Answer Prediction runs will be evaluated using standard measures for evaluation of prediction tasks with AUC being the primary measure.

## Runs 

Participating groups will be allowed to submit as many runs as they like, but they need authorization from the Track organizers before submitting more than 10 runs per task. Not all runs are likely to be used for pooling and groups will need to specify a preference ordering for pooling purposes.

Runs may be either automatic or manual. 

*Automatic runs:* Only the topic's `query` or `question` field may be used for automatic runs.  An automatic run should only use the `query` or the `question` field, but not both. An automatic run is made without any tuning or customization based on the topics.  Best practice for an automatic run is to avoid using the topics or even looking at them until all decisions and code have been written to produce the automatic run. 

*Manual runs:* A manual run is anything that is not an automatic run. Manual runs commonly have some human input based on the topics, e.g., hand-crafted queries or relevance feedback. All topic fields may be used for manual runs.  We encourage manual runs in addition to automatic runs.  

Submission format for Web Retrieval runs will follow the standard TREC run format.  For each topic, please return 1,000 ranked documents. The standard TREC run format is as follows:

```
qid Q0 docno rank score tag
```
where:
* `qid`: the topic number;
* `Q0`: unused and should always be Q0;
* `docno`: the document id number returned by your system for the topic `qid` (See above for more details on how docnos should be constructed it); 
* `rank`: the rank the document is retrieved;
* `score`: the score (integer or floating point) that generated the ranking. The score must be in descending (non-increasing) order. The score is important to handle tied scores. (`trec_eval` sorts documents by the specified scores values and not your ranks values);
* `runtag`: a tag that uniquely identifies your group AND the method you used to produce the run. Each run should have a different tag.  **Runtags for runs submitted by one group should all share a common prefix to identify the group across runs.**

The fields should be separated with a space. 

An example run is shown below:
```
1 Q0 en.noclean.c4-train.04124-of-07168.69102 1 14.8928003311 myGroupNameMyMethodName
1 Q0 en.noclean.c4-train.03346-of-07168.52165 2 14.7590999603 myGroupNameMyMethodName
1 Q0 en.noclean.c4-train.03904-of-07168.54203 3 14.5707998276 myGroupNameMyMethodName
...
```

The submission format for the Answer Prediction task will be four columns, each separated by a space:
```
qid answer score runtag
```
where:
+ `qid`: the topic number
+ `answer`: either the string "yes" or "no"
+ `score`: A floating point value ranging from 1.0 to 0.0 where 1.0 means "yes" and a score of 0.0 means "no".  The scores should be comparable across topics.
+ `runtag`: a tag that uniquely identifies your group AND the method you used to produce the run. Each run should have a different tag.  **Runtags for runs submitted by one group should all share a common prefix to identify the group across runs.**

An example run is shown below:
```
151 yes 0.95234 myGroupNameMyMethodName
152 no 0.30218 myGroupNameMyMethodName
153 no 0.00396 myGroupNameMyMethodName
...
```

## Schedule
* **June 21, 2021** Collection released;
* **May 12, 2022** Guidelines finalized;
* **July 13, 2022** Topics released via [TREC's web page for active participants](https://trec.nist.gov/act_part/tracks2022.html)
* ~~August 1, 2022~~ **August 28, 2022** Runs due;
* **Tentative: End of September 2022** Results returned;
* **Tentative: October 2022** Notebook paper due;
* **November 14-18, 2022** TREC Conference;
* **February 2023** Final report due.

## Organizers

* [Charles Clarke, University of Waterloo](https://cs.uwaterloo.ca/about/people/claclark)
* [Maria Maistro, University of Copenhagen](https://di.ku.dk/english/staff/?pure=en/persons/641366)
* [Mark Smucker, University of Waterloo](https://uwaterloo.ca/management-sciences/profile/msmucker)

