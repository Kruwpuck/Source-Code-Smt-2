import numpy as np
import pandas as pd
from sklearn.model_selection import train_test_split
from sklearn.naive_bayes import GaussianNB
from sklearn.neighbors import KNeighborsClassifier
from sklearn.metrics import classification_report, accuracy_score, confusion_matrix

# Membuat dataset contoh
np.random.seed(0)
data_size = 1000
features = 10

# 80% positif (1), 20% negatif (0)
labels = np.random.choice([0, 1], size=data_size, p=[0.2, 0.8])
data = np.random.rand(data_size, features)
dataset = pd.DataFrame(data, columns=[f'feature_{i}' for i in range(features)])
dataset['label'] = labels

# Membagi data menjadi set pelatihan dan pengujian
X = dataset.drop('label', axis=1)
y = dataset['label']
X_train, X_test, y_train, y_test = train_test_split(X, y, test_size=0.2, random_state=0)

# Naïve Bayes
nb_model = GaussianNB()
nb_model.fit(X_train, y_train)
y_pred_nb = nb_model.predict(X_test)

print("Naïve Bayes Classification Report:")
print(classification_report(y_test, y_pred_nb))
print("Naïve Bayes Accuracy:", accuracy_score(y_test, y_pred_nb))
print("Naïve Bayes Confusion Matrix:\n", confusion_matrix(y_test, y_pred_nb))

# k-Nearest Neighbors
k = 5  # k dapat di-tuning untuk performa terbaik
knn_model = KNeighborsClassifier(n_neighbors=k)
knn_model.fit(X_train, y_train)
y_pred_knn = knn_model.predict(X_test)

print("\nk-Nearest Neighbors Classification Report:")
print(classification_report(y_test, y_pred_knn))
print("k-Nearest Neighbors Accuracy:", accuracy_score(y_test, y_pred_knn))
print("k-Nearest Neighbors Confusion Matrix:\n", confusion_matrix(y_test, y_pred_knn))

# Membandingkan hasil
nb_accuracy = accuracy_score(y_test, y_pred_nb)
knn_accuracy = accuracy_score(y_test, y_pred_knn)

print(f"\nComparison of Model Performance:\nNaïve Bayes Accuracy: {nb_accuracy}\nk-Nearest Neighbors Accuracy: {knn_accuracy}")

if nb_accuracy > knn_accuracy:
    print("Naïve Bayes performs better for this dataset.")
else:
    print("k-Nearest Neighbors performs better for this dataset.")
