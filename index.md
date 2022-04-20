---
layout: default
---

# TREC Health Misinformation Track (2022) - Draft 

**WARNING: These guidelines are not yet finalized.  These are draft guidelines.**

## Track Introduction

Web search engines are frequently used to help people make decisions about health-related issues.  Unfortunately, the web is filled with misinformation regarding the efficacy of treatments for health issues.  Search users may not be able to discern correct from incorrect information, nor credible from non-credible sources.  As a result of finding misinformation deemed by the user to be useful to their decision making task, they can make incorrect decisions that waste money and put their health at risk.

The TREC Health Misinformation track fosters research on retrieval methods that promote reliable and correct information over misinformation for health-related decision making tasks.

### Track communication

Track announcements are made via the #health-misinfo-2022 channel in the [TREC Slack](https://trectalk.slack.com).  

## 2022 Track Tasks

The 2022 track will have two tasks.  We will again repeat the AdHoc Web Retrieval task, and we will add a Stance Prediction task.

### AdHoc Web Retrieval

**Task Description:** Participants devise search technologies that promote credible and correct information over incorrect information, with the assumption that correct information can better lead people to make correct decisions.

Each topic concerns itself with a health issue and a treatment for that issue.  The topics represent a user who is looking for information that is useful for making a decision about whether or not the treatment is helpful or unhelpful for treating the health issue.  Better search result rankings will place very useful documents that are credible and correct at the top of ranking, and will not return incorrect information.  In this search task, incorrect information is considered harmful and should not be returned.  

For each topic, we have chosen a 'stance' for the topic on whether the treatment is helpful or unhelpful for the health issue.  **We do not claim to be providing medical advice, and medical decisions should never be made based on the stance we have chosen.  Consult a medical doctor for professional advice.**  If a treatment is considered 'helpful', then correct documents will those construed to be supportive of the treatment and incorrect documents will be those that would dissuade the searcher from the treatment.  Likewise, an 'unhelpful' treatment should return documents that dissuade the searcher from using the treatment and should avoid returning documents that are supportive of using the treatment.  For each topic, the topic's author will determine an `evidence` link for a webpage that the topic author used as the basis for the topic's stance.  

In 2022, the topic's `stance` and `evidence` fields will not be revealed until after evaluation results are provided by NIST.  

The primary ad-hoc task is to use the topic's `query` or `description` for a topic and return a ranking of documents without use of any of the other topic fields.  

Runs may be either automatic runs or manual runs.  Automatic runs are required to make no use of the provided topics except for final production of the run.  Automatic runs should use only the `query` or the `description` field, but not both.  Submitted automatic runs will be required to be designated as `query` or `description` runs.  

The 2022 topics will be similar to the 2021 topics, and if needed for an automatic run, the 2021 topics should be used for tuning and not the 2022 topics.

Manual runs are any runs that are not automatic runs.  A manual run typically has been tuned on the topics or hand human intervention to improve performance.  For example, a human could manually determine a topic's stance and feed that judgment to a ranking method, and thus the run would be a manual run.  Any human rewriting of the query or description fields would also make a run be a manual run.

### Stance Prediction

As noted above, in 2022, we will not provide a topic's stance until after evaluation.  In 2020 and 2021, the use of a topic's stance has been important to the success of many submitted runs.  The Stance Prediction task provide a chance for participants to focus on the challenge of predicting the stance of a topic.  

For each topic, participants will predict a topic's stance as either "helpful" or "unhelpful".  Participants will also provide a prediction score between 0 and 1 for each topic.  A score of 1 means "helpful" and a score of 0 means "unhelpful".  The prediction scores will allow the computation of AUC as part of the evaluation of the prediction quality.

As with the AdHoc task, Stance Prediction runs may be automatic or manual and will follow the same rules.  See above for details.

#### Document Collection

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

#### Topics

Topics will be authored by the track organizers.  The NIST assessors will be provided the topic's query, description, and narrative and be asked to make judgments as per the assessing guidelines.  The 2022 guidelines are to be written, but they will be similar to the [2021 assessing guidelines](docs/TREC-2021-Health-Misinformation-Track-Assessing-Guidelines_Version-2.pdf).

The topics will be provided as XML files using the following format:

```xml
<topics>
  <topic>
    <number>1234</number>
    <query>dexamethasone croup</query>
    <description>Is dexamethasone a good treatment for croup?</description>
    <narrative>Croup is an infection of the upper airway and causes swelling, 
      which obstructs breathing and leads to a barking cough. As one kind of 
      corticosteroids, dexamethasone can weaken the immune response and 
      therefore mitigate symptoms such as swelling. A very useful document 
      would discuss the effectiveness of dexamethasone for croup, i.e. a very 
      useful document specifically addresses or answers the search topic's 
      question. A useful document would provide information that would help 
      a user make a decision about treating croup with dexamethasone, and 
      may discuss either separately or jointly: croup, recommended treatments 
      for croup, the pros and cons of dexamethasone, etc.</narrative>
    <disclaimer>We do not claim to be providing medical advice, and medical 
      decisions should never be made based on the stance we have chosen.  
      Consult a medical doctor for professional advice.</disclaimer>
  </topic>
<topic>
...
</topic>
</topics>
```

#### Evaluation

The evaluation of AdHoc runs will be similar to 2021 but with likely improvements by the organizers.  The Stance Prediction runs will be evaluated using standard measures for evaluation of prediction tasks with AUC being the primary measure.

#### Runs 

Participating groups will be allowed to submit as many runs as they like, but they need authorization from the Track organizers before submitting more than 10 runs. Not all runs are likely to be used for pooling and groups will need to specify a preference ordering for pooling purposes.

Runs may be either automatic or manual. 

*Automatic runs:* Only the topic's `query` or `description` field may be used for automatic runs.  An automatic run may only use the `query` or the `description` field, but not both. An automatic run is made without any tuning or customization based on the topics.  Best practice for an automatic run is to avoid using the topics or even looking at them until all decisions and code have been written to produce the automatic run. 

*Manual runs:* A manual run is anything that is not an automatic run. Manual runs commonly have some human input based on the topics, e.g., hand-crafted queries or relevance feedback. All topic fields may be used for manual runs.  We encourage manual runs in addition to automatic runs.  

Submission format for AdHoc runs will follow the standard TREC run format.  For each topic, please return 1,000 ranked documents. The standard TREC run format is as follows:

```
qid Q0 docno rank score tag
```
where:
* `qid`: the topic number;
* `Q0`: unused and should always be Q0;
* `docno`: the document id number returned by your system for the topic `qid` (See above for more details on how docnos should be constructed it); 
* `ran`: the rank the document is retrieved;
* `score`: the score (integer or floating point) that generated the ranking. The score must be in descending (non-increasing) order. The score is important to handle tied scores. (`trec_eval` sorts documents by the specified scores values and not your ranks values);
* `tag`: a tag that uniquely identifies your group AND the method you used to produce the run. Each run should have a different tag.

The fields should be separated with a space. 

An example run is shown below:
```
1 Q0 en.noclean.c4-train.04124-of-07168.69102 1 14.8928003311 myGroupNameMyMethodName
1 Q0 en.noclean.c4-train.03346-of-07168.52165 2 14.7590999603 myGroupNameMyMethodName
1 Q0 en.noclean.c4-train.03904-of-07168.54203 3 14.5707998276 myGroupNameMyMethodName
...
```

The submission format for the Stance Prediction task will be forthcoming.

## Schedule
* **June 21, 2021** Collection released;
* **Tentative: July 2022** Topics released;
* **Tentative: September, 2022** Runs due;
* **Tentative: End of September 2022** Results returned;
* **Tentative: October 2022** Notebook paper due;
* **November 14-18, 2022** TREC Conference;
* **February 2023** Final report due.

## Organizers

* [Charles Clarke, University of Waterloo](https://cs.uwaterloo.ca/about/people/claclark)
* [Maria Maistro, University of Copenhagen](https://di.ku.dk/english/staff/?pure=en/persons/641366)
* [Mark Smucker, University of Waterloo](http://mansci.uwaterloo.ca/~msmucker/)

