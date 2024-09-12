from fastapi import FastAPI
from pydantic import BaseModel
import tensorflow as tf

app = FastAPI()

# 假设你的 TextCNN 模型保存在 model/ 目录下
model = tf.keras.models.load_model('./model/textcnn_model.h5')

# 定义输入和输出的数据模型
class Memthing(BaseModel):
    content: str

class ClassificationResult(BaseModel):
    classification: str
    grade: float

# 定义分类路由
@app.post("/classify", response_model=ClassificationResult)
async def classify(memthing: Memthing):
    # 将 memthing.content 预处理为模型可以接受的格式
    input_data = preprocess(memthing.content)
    prediction = model.predict(input_data)

    # 解析分类结果
    classification = decode_classification(prediction)
    grade = calculate_grade(prediction)

    return ClassificationResult(classification=classification, grade=grade)

def preprocess(text):
    """
    预处理文本，转换为模型可以接受的输入格式。
    """
    # 将文本转换为词索引
    sequences = tokenizer.texts_to_sequences([text])
    # 将序列填充或截断为固定长度
    padded_sequences = pad_sequences(sequences, maxlen=MAX_SEQUENCE_LENGTH)
    return padded_sequences

def decode_classification(prediction):
    """
    解码模型的预测结果，将预测结果映射为类别标签。
    """
    # 模型输出是概率分布，找到概率最高的类别索引
    predicted_class_index = prediction.argmax(axis=1)[0]
    # 返回对应的类别标签
    return CLASS_LABELS[predicted_class_index]

def calculate_grade(prediction):
    """
    根据模型输出的预测概率计算评分。
    """
    # 假设模型输出的是一个概率分布，选择最高的概率
    max_probability = prediction.max()
    # 将概率转换为 0 到 100 之间的评分
    grade = max_probability * 100
    return grade